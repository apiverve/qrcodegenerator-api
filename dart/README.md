# QR Code Generator API - Dart/Flutter Client

QR Code Generator creates customizable QR codes with support for colors, gradients, logos, and various styling options. Generate professional QR codes for marketing, packaging, and digital experiences.

[![pub package](https://img.shields.io/pub/v/apiverve_qrcodegenerator.svg)](https://pub.dev/packages/apiverve_qrcodegenerator)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

This is the Dart/Flutter client for the [QR Code Generator API](https://apiverve.com/marketplace/qrcodegenerator?utm_source=dart&utm_medium=referral&utm_content=readme).

## Installation

Add this to your `pubspec.yaml`:

```yaml
dependencies:
  apiverve_qrcodegenerator: ^1.2.6
```

Then run:

```bash
dart pub get
# or for Flutter
flutter pub get
```

## Usage

```dart
import 'package:apiverve_qrcodegenerator/apiverve_qrcodegenerator.dart';

void main() async {
  final client = QrcodegeneratorClient('YOUR_API_KEY');

  try {
    final response = await client.execute({
      'value': 'https://apiverve.com',
      'type': 'url',
      'format': 'png',
      'margin': 0
    });

    print('Status: ${response.status}');
    print('Data: ${response.data}');
  } catch (e) {
    print('Error: $e');
  }
}
```

## Response

```json
{
  "status": "ok",
  "error": null,
  "data": {
    "id": "b07577c8-e17f-4af3-aeff-94c74d9ea04a",
    "format": "png",
    "type": "url",
    "correction": "M",
    "size": 5,
    "margin": 0,
    "expires": 1766096864874,
    "downloadURL": "https://storage.googleapis.com/apiverve/APIData/qrcodegenerator/2556d708-7ec4-45ce-b0fb-98d47fcf170f.png?GoogleAccessId=635500398038-compute%40developer.gserviceaccount.com&Expires=1766096864&Signature=fRvOFggSHakWfnBqT46dyv9trO4sVcwPjvjV8hpFG9yI2G2UG%2FoAvJFwzX%2F%2FYPcxRfzcR5Cnc6uSx74iJANJXn42Hpcgr5EudNa%2BxtAo106LfyWBaQgaZggNVai4D6WsDYyBL8sTPGIDxjnldVFIJy7HFMRCKUZ%2B3KQMiYujmFoc2tVMVi7lngxrfAFIylHOmEgCXDaAWXGb6NwV%2BZNsDE1Ju9DLseLnhZ%2Fh2LxIbxRiF042MLkg2hp54P%2FrIhABfOcwW7CoV9Hpn4N88t5818ljyCrLfeLQpxM5gtcRCZJ3MYiICwLkmr%2F62KDmx5JBMFLB%2FVfh%2Bewsc0QJytB3vg%3D%3D"
  }
}
```

## API Reference

- **API Home:** [QR Code Generator API](https://apiverve.com/marketplace/qrcodegenerator?utm_source=dart&utm_medium=referral&utm_content=readme)
- **Documentation:** [docs.apiverve.com/ref/qrcodegenerator](https://docs.apiverve.com/ref/qrcodegenerator?utm_source=dart&utm_medium=referral&utm_content=readme)

## Authentication

All requests require an API key. Get yours at [apiverve.com](https://apiverve.com?utm_source=dart&utm_medium=referral&utm_content=readme).

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with Dart for [APIVerve](https://apiverve.com?utm_source=dart&utm_medium=referral&utm_content=readme)
