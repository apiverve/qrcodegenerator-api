using System;
using System.Collections.Generic;
using System.Text;
using Newtonsoft.Json;

namespace APIVerve.API.QRCodeGenerator
{
    /// <summary>
    /// Query options for the QR Code Generator API
    /// </summary>
    public class QRCodeGeneratorQueryOptions
    {
        /// <summary>
        /// The text or data to encode in the QR code
        /// </summary>
        [JsonProperty("value")]
        public string Value { get; set; }

        /// <summary>
        /// The type of data being encoded. Advanced types (wifi, vcard) are premium.
        /// </summary>
        [JsonProperty("type")]
        public string Type { get; set; }

        /// <summary>
        /// Output format. Vector formats (svg, webp) are premium.
        /// </summary>
        [JsonProperty("format")]
        public string Format { get; set; }

        /// <summary>
        /// Size of the QR code in pixels (50-2048)
        /// </summary>
        [JsonProperty("size")]
        public double? Size { get; set; }

        /// <summary>
        /// Margin around the QR code in pixels (0-100)
        /// </summary>
        [JsonProperty("margin")]
        public double? Margin { get; set; }

        /// <summary>
        /// Foreground color as hex code (e.g., #000000)
        /// </summary>
        [JsonProperty("color")]
        public string Color { get; set; }

        /// <summary>
        /// Background color as hex code (e.g., #ffffff)
        /// </summary>
        [JsonProperty("backgroundColor")]
        public string BackgroundColor { get; set; }

        /// <summary>
        /// Style of QR code dots
        /// </summary>
        [JsonProperty("dotStyle")]
        public string DotStyle { get; set; }

        /// <summary>
        /// Style of corner squares
        /// </summary>
        [JsonProperty("cornerSquareStyle")]
        public string CornerSquareStyle { get; set; }

        /// <summary>
        /// Style of corner dots
        /// </summary>
        [JsonProperty("cornerDotStyle")]
        public string CornerDotStyle { get; set; }

        /// <summary>
        /// Gradient configuration with type (linear, radial) and colorStops array
        /// </summary>
        [JsonProperty("gradient")]
        public object? Gradient { get; set; }

        /// <summary>
        /// URL of logo image to place in center of QR code
        /// </summary>
        [JsonProperty("logo")]
        public string Logo { get; set; }

        /// <summary>
        /// Size of logo relative to QR code (0.1-0.5)
        /// </summary>
        [JsonProperty("logoSize")]
        public double? LogoSize { get; set; }

        /// <summary>
        /// Margin around logo in pixels
        /// </summary>
        [JsonProperty("logoMargin")]
        public double? LogoMargin { get; set; }
    }
}
