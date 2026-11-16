import 'dart:convert';
import 'package:http/http.dart' as http;
import 'models.dart';

/// Validation rule for a parameter.
class ValidationRule {
  final String type;
  final bool required;
  final num? min;
  final num? max;
  final int? minLength;
  final int? maxLength;
  final String? format;
  final List<String>? enumValues;

  const ValidationRule({
    required this.type,
    required this.required,
    this.min,
    this.max,
    this.minLength,
    this.maxLength,
    this.format,
    this.enumValues,
  });
}

/// Exception thrown when parameter validation fails.
class QrcodegeneratorValidationException implements Exception {
  final List<String> errors;

  QrcodegeneratorValidationException(this.errors);

  @override
  String toString() => 'QrcodegeneratorValidationException: ${errors.join("; ")}';
}

/// Format validation patterns.
final _formatPatterns = {
  'email': RegExp(r'^[^\s@]+@[^\s@]+\.[^\s@]+$'),
  'url': RegExp(r'^https?://.+'),
  'ip': RegExp(r'^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$|^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$'),
  'date': RegExp(r'^\d{4}-\d{2}-\d{2}$'),
  'hexColor': RegExp(r'^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$'),
};

/// QR Code Generator API client.
///
/// For more information, visit: https://apiverve.com/marketplace/qrcodegenerator?utm_source=dart&utm_medium=referral&utm_content=readme
///
/// Parameters:
/// * [value] (required) - The text or data to encode in the QR code
/// * [type] - The type of data being encoded. Advanced types (wifi, vcard) are premium.
/// * [format] - Output format. Vector formats (svg, webp) are premium.
/// * [size] - Size of the QR code in pixels (50-2048)
/// * [margin] - Margin around the QR code in pixels (0-100)
/// * [color] - Foreground color as hex code (e.g., #000000)
/// * [backgroundColor] - Background color as hex code (e.g., #ffffff)
/// * [dotStyle] - Style of QR code dots
/// * [cornerSquareStyle] - Style of corner squares
/// * [cornerDotStyle] - Style of corner dots
/// * [gradient] - Gradient configuration with type (linear, radial) and colorStops array
/// * [logo] - URL of logo image to place in center of QR code [format: url]
/// * [logoSize] - Size of logo relative to QR code (0.1-0.5)
/// * [logoMargin] - Margin around logo in pixels
class QrcodegeneratorClient {
  final String apiKey;
  final String baseUrl;
  final http.Client _httpClient;

  /// Validation rules for parameters.
  static final Map<String, ValidationRule> _validationRules = <String, ValidationRule>{
    'value': ValidationRule(type: 'string', required: true),
    'type': ValidationRule(type: 'string', required: false),
    'format': ValidationRule(type: 'string', required: false),
    'size': ValidationRule(type: 'number', required: false),
    'margin': ValidationRule(type: 'number', required: false),
    'color': ValidationRule(type: 'string', required: false),
    'backgroundColor': ValidationRule(type: 'string', required: false),
    'dotStyle': ValidationRule(type: 'string', required: false),
    'cornerSquareStyle': ValidationRule(type: 'string', required: false),
    'cornerDotStyle': ValidationRule(type: 'string', required: false),
    'gradient': ValidationRule(type: 'object', required: false),
    'logo': ValidationRule(type: 'string', required: false, format: 'url'),
    'logoSize': ValidationRule(type: 'number', required: false),
    'logoMargin': ValidationRule(type: 'number', required: false),
  };

  QrcodegeneratorClient(this.apiKey, {
    this.baseUrl = 'https://api.apiverve.com/v1/qrcodegenerator',
    http.Client? httpClient,
  }) : _httpClient = httpClient ?? http.Client();

  /// Validates parameters against defined rules.
  /// Throws [QrcodegeneratorValidationException] if validation fails.
  void _validateParams(Map<String, dynamic> params) {
    final errors = <String>[];

    for (final entry in _validationRules.entries) {
      final paramName = entry.key;
      final rule = entry.value;
      final value = params[paramName];

      // Check required
      if (rule.required && (value == null || (value is String && value.isEmpty))) {
        errors.add('Required parameter [$paramName] is missing');
        continue;
      }

      if (value == null) continue;

      // Type-specific validation
      if (rule.type == 'integer' || rule.type == 'number') {
        final numValue = value is num ? value : num.tryParse(value.toString());
        if (numValue == null) {
          errors.add('Parameter [$paramName] must be a valid ${rule.type}');
          continue;
        }
        if (rule.min != null && numValue < rule.min!) {
          errors.add('Parameter [$paramName] must be at least ${rule.min}');
        }
        if (rule.max != null && numValue > rule.max!) {
          errors.add('Parameter [$paramName] must be at most ${rule.max}');
        }
      } else if (rule.type == 'string' && value is String) {
        if (rule.minLength != null && value.length < rule.minLength!) {
          errors.add('Parameter [$paramName] must be at least ${rule.minLength} characters');
        }
        if (rule.maxLength != null && value.length > rule.maxLength!) {
          errors.add('Parameter [$paramName] must be at most ${rule.maxLength} characters');
        }
        if (rule.format != null && _formatPatterns.containsKey(rule.format)) {
          if (!_formatPatterns[rule.format]!.hasMatch(value)) {
            errors.add('Parameter [$paramName] must be a valid ${rule.format}');
          }
        }
      }

      // Enum validation
      if (rule.enumValues != null && rule.enumValues!.isNotEmpty) {
        if (!rule.enumValues!.contains(value.toString())) {
          errors.add('Parameter [$paramName] must be one of: ${rule.enumValues!.join(", ")}');
        }
      }
    }

    if (errors.isNotEmpty) {
      throw QrcodegeneratorValidationException(errors);
    }
  }

  /// Execute a request to the QR Code Generator API.
  ///
  /// Parameters are validated before sending the request.
  Future<QrcodegeneratorResponse> execute(Map<String, dynamic> params) async {
    // Validate parameters
    _validateParams(params);
    if (apiKey.isEmpty) {
      throw QrcodegeneratorException('API key is required. Get your API key at: https://apiverve.com');
    }

    try {
      final response = await _httpClient.post(
        Uri.parse(baseUrl),
        headers: {
          'x-api-key': apiKey,
          'Content-Type': 'application/json',
        },
        body: jsonEncode(params),
      );

      if (response.statusCode == 200) {
        final json = jsonDecode(response.body) as Map<String, dynamic>;
        return QrcodegeneratorResponse.fromJson(json);
      } else if (response.statusCode == 401) {
        throw QrcodegeneratorException('Invalid API key');
      } else if (response.statusCode == 404) {
        throw QrcodegeneratorException('Resource not found');
      } else {
        throw QrcodegeneratorException('API error: ${response.statusCode}');
      }
    } catch (e) {
      if (e is QrcodegeneratorException) rethrow;
      throw QrcodegeneratorException('Request failed: $e');
    }
  }


  /// Close the HTTP client.
  void close() {
    _httpClient.close();
  }
}

/// Exception thrown by the QR Code Generator API client.
class QrcodegeneratorException implements Exception {
  final String message;

  QrcodegeneratorException(this.message);

  @override
  String toString() => 'QrcodegeneratorException: $message';
}
