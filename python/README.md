QR Code Generator API
============

QR Code Generator creates customizable QR codes with support for colors, gradients, logos, and various styling options. Generate professional QR codes for marketing, packaging, and digital experiences.


[![PyPI version](https://img.shields.io/pypi/v/apiverve-qrcodegenerator.svg)](https://pypi.org/project/apiverve-qrcodegenerator/)

This is a Python wrapper for the [QR Code Generator API](https://apiverve.com/marketplace/qrcodegenerator?utm_source=pypi&utm_medium=referral&utm_content=readme)

---

## Installation

Using `pip`:

```bash
pip install apiverve-qrcodegenerator
```

Using `pip3`:

```bash
pip3 install apiverve-qrcodegenerator
```

---

## Configuration

Before using the QR Code Generator API client, you have to setup your account and obtain your API Key.
You can get it by signing up at [https://apiverve.com](https://apiverve.com?utm_source=pypi&utm_medium=referral&utm_content=readme)

---

## Quick Start

Here's a simple example to get you started quickly:

```python
from apiverve_qrcodegenerator.apiClient import QrcodegeneratorAPIClient

# Initialize the client with your APIVerve API key
api = QrcodegeneratorAPIClient("[YOUR_API_KEY]")

query = {
    "value": "https://apiverve.com",
    "type": "url",
    "format": "png",
    "margin": 0
}

try:
    # Make the API call
    result = api.execute(query)

    # Print the result
    print(result)
except Exception as e:
    print(f"Error: {e}")
```

---

## Usage

The QR Code Generator API documentation is found here: [https://docs.apiverve.com/ref/qrcodegenerator](https://docs.apiverve.com/ref/qrcodegenerator?utm_source=pypi&utm_medium=referral&utm_content=readme).
You can find parameters, example responses, and status codes documented here.

### Setup

```python
# Import the client module
from apiverve_qrcodegenerator.apiClient import QrcodegeneratorAPIClient

# Initialize the client with your APIVerve API key
api = QrcodegeneratorAPIClient("[YOUR_API_KEY]")
```

---

## Perform Request

Using the API client, you can perform requests to the API.

###### Define Query

```python
query = {
    "value": "https://apiverve.com",
    "type": "url",
    "format": "png",
    "margin": 0
}
```

###### Simple Request

```python
# Make a request to the API
result = api.execute(query)

# Print the result
print(result)
```

###### Example Response

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

---

## Error Handling

The API client provides comprehensive error handling through the `QrcodegeneratorAPIClientError` exception. Here are some examples:

### Basic Error Handling

```python
from apiverve_qrcodegenerator.apiClient import QrcodegeneratorAPIClient, QrcodegeneratorAPIClientError

api = QrcodegeneratorAPIClient("[YOUR_API_KEY]")

query = {
    "value": "https://apiverve.com",
    "type": "url",
    "format": "png",
    "margin": 0
}

try:
    result = api.execute(query)
    print("Success!")
    print(result)
except QrcodegeneratorAPIClientError as e:
    print(f"API Error: {e.message}")
    if e.status_code:
        print(f"Status Code: {e.status_code}")
    if e.response:
        print(f"Response: {e.response}")
```

### Handling Specific Error Types

```python
from apiverve_qrcodegenerator.apiClient import QrcodegeneratorAPIClient, QrcodegeneratorAPIClientError

api = QrcodegeneratorAPIClient("[YOUR_API_KEY]")

query = {
    "value": "https://apiverve.com",
    "type": "url",
    "format": "png",
    "margin": 0
}

try:
    result = api.execute(query)

    # Check for successful response
    if result.get('status') == 'success':
        print("Request successful!")
        print(result.get('data'))
    else:
        print(f"API returned an error: {result.get('error')}")

except QrcodegeneratorAPIClientError as e:
    # Handle API client errors
    if e.status_code == 401:
        print("Unauthorized: Invalid API key")
    elif e.status_code == 429:
        print("Rate limit exceeded")
    elif e.status_code >= 500:
        print("Server error - please try again later")
    else:
        print(f"API error: {e.message}")
except Exception as e:
    # Handle unexpected errors
    print(f"Unexpected error: {str(e)}")
```

### Using Context Manager (Recommended)

The client supports the context manager protocol for automatic resource cleanup:

```python
from apiverve_qrcodegenerator.apiClient import QrcodegeneratorAPIClient, QrcodegeneratorAPIClientError

query = {
    "value": "https://apiverve.com",
    "type": "url",
    "format": "png",
    "margin": 0
}

# Using context manager ensures proper cleanup
with QrcodegeneratorAPIClient("[YOUR_API_KEY]") as api:
    try:
        result = api.execute(query)
        print(result)
    except QrcodegeneratorAPIClientError as e:
        print(f"Error: {e.message}")
# Session is automatically closed here
```

---

## Advanced Features

### Debug Mode

Enable debug logging to see detailed request and response information:

```python
from apiverve_qrcodegenerator.apiClient import QrcodegeneratorAPIClient

# Enable debug mode
api = QrcodegeneratorAPIClient("[YOUR_API_KEY]", debug=True)

query = {
    "value": "https://apiverve.com",
    "type": "url",
    "format": "png",
    "margin": 0
}

# Debug information will be printed to console
result = api.execute(query)
```

### Manual Session Management

If you need to manually manage the session lifecycle:

```python
from apiverve_qrcodegenerator.apiClient import QrcodegeneratorAPIClient

api = QrcodegeneratorAPIClient("[YOUR_API_KEY]")

query = {
    "value": "https://apiverve.com",
    "type": "url",
    "format": "png",
    "margin": 0
}

try:
    result = api.execute(query)
    print(result)
finally:
    # Manually close the session when done
    api.close()
```

---

## Customer Support

Need any assistance? [Get in touch with Customer Support](https://apiverve.com/contact?utm_source=pypi&utm_medium=referral&utm_content=readme).

---

## Updates
Stay up to date by following [@apiverveHQ](https://twitter.com/apiverveHQ) on Twitter.

---

## Legal

All usage of the APIVerve website, API, and services is subject to the [APIVerve Terms of Service](https://apiverve.com/terms?utm_source=pypi&utm_medium=referral&utm_content=readme) and all legal documents and agreements.

---

## License
Licensed under the MIT License (MIT)

Copyright (&copy;) 2026 APIVerve, and EvlarSoft LLC

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
