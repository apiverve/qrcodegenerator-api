# QR Code Generator API - Ruby Gem

QR Code Generator creates customizable QR codes with support for colors, gradients, logos, and various styling options. Generate professional QR codes for marketing, packaging, and digital experiences.

## Installation

Add this line to your application's Gemfile:

```ruby
gem 'apiverve_qrcodegenerator'
```

And then execute:

```bash
$ bundle install
```

Or install it yourself as:

```bash
$ gem install apiverve_qrcodegenerator
```

## Getting Started

Get your API key at [APIVerve](https://apiverve.com)

### Basic Usage

```ruby
require 'apiverve_qrcodegenerator'

# Initialize the client
client = APIVerve::Qrcodegenerator::Client.new(api_key: "YOUR_API_KEY")

# Make a request
response = client.execute({
  value: "https://apiverve.com",
  type: "url",
  format: "png",
  margin: 0
})

# Print the response
puts response
```


### Error Handling

```ruby
begin
  response = client.execute({ value: "https://apiverve.com", type: "url", format: "png", margin: 0 })
  puts response["data"]
rescue APIVerve::Qrcodegenerator::ValidationError => e
  puts "Validation error: #{e.errors.join(', ')}"
rescue APIVerve::Qrcodegenerator::APIError => e
  puts "API error: #{e.message}"
  puts "Status code: #{e.status_code}"
end
```

### Debug Mode

```ruby
# Enable debug logging
client = APIVerve::Qrcodegenerator::Client.new(
  api_key: "YOUR_API_KEY",
  debug: true
)
```

## Example Response

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

## Documentation

For more information, visit the [API Documentation](https://docs.apiverve.com/ref/qrcodegenerator?utm_source=rubygems&utm_medium=referral&utm_content=readme).

## Support

- Website: [https://apiverve.com/marketplace/qrcodegenerator?utm_source=ruby&utm_medium=referral&utm_content=readme](https://apiverve.com/marketplace/qrcodegenerator?utm_source=ruby&utm_medium=referral&utm_content=readme)
- Email: hello@apiverve.com

## License

This gem is available under the [MIT License](LICENSE).
