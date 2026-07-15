"""
QR Code Generator API - Basic Usage Example

This example demonstrates the basic usage of the QR Code Generator API.
API Documentation: https://docs.apiverve.com/ref/qrcodegenerator
"""

import os
import requests
import json

API_KEY = os.getenv('APIVERVE_API_KEY', 'YOUR_API_KEY_HERE')
API_URL = 'https://api.apiverve.com/v1/qrcodegenerator'

def call_qrcodegenerator_api():
    """
    Make a POST request to the QR Code Generator API
    """
    try:
        # Request body
        request_body &#x3D; {
    &#x27;value&#x27;: &#x27;https://apiverve.com&#x27;,
    &#x27;type&#x27;: &#x27;url&#x27;,
    &#x27;format&#x27;: &#x27;png&#x27;,
    &#x27;margin&#x27;: &#x27;0&#x27;
}

        headers = {
            'x-api-key': API_KEY,
            'Content-Type': 'application/json'
        }

        response = requests.post(API_URL, headers=headers, json=request_body)

        # Raise exception for HTTP errors
        response.raise_for_status()

        data = response.json()

        # Check API response status
        if data.get('status') == 'ok':
            print('✓ Success!')
            print('Response data:', json.dumps(data['data'], indent=2))
            return data['data']
        else:
            print('✗ API Error:', data.get('error', 'Unknown error'))
            return None

    except requests.exceptions.RequestException as e:
        print(f'✗ Request failed: {e}')
        return None

if __name__ == '__main__':
    print('📤 Calling QR Code Generator API...\n')

    result = call_qrcodegenerator_api()

    if result:
        print('\n📊 Final Result:')
        print(json.dumps(result, indent=2))
    else:
        print('\n✗ API call failed')
