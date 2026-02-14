# QR Code Generator API

QR Code Generator creates customizable QR codes with support for colors, gradients, logos, and various styling options. Generate professional QR codes for marketing, packaging, and digital experiences.

![Build Status](https://img.shields.io/badge/build-passing-green)
![Code Climate](https://img.shields.io/badge/maintainability-B-purple)
![Prod Ready](https://img.shields.io/badge/production-ready-blue)
[![npm version](https://img.shields.io/npm/v/@apiverve/qrcodegenerator.svg)](https://www.npmjs.com/package/@apiverve/qrcodegenerator)

This is a Javascript Wrapper for the [QR Code Generator API](https://apiverve.com/marketplace/qrcodegenerator?utm_source=npm&utm_medium=readme)

---

## Installation

Using npm:
```shell
npm install @apiverve/qrcodegenerator
```

Using yarn:
```shell
yarn add @apiverve/qrcodegenerator
```

---

## Configuration

Before using the QR Code Generator API client, you have to setup your account and obtain your API Key.
You can get it by signing up at [https://apiverve.com](https://apiverve.com?utm_source=npm&utm_medium=readme)

---

## Quick Start

[Get started with the Quick Start Guide](https://docs.apiverve.com/quickstart?utm_source=npm&utm_medium=readme)

The QR Code Generator API documentation is found here: [https://docs.apiverve.com/ref/qrcodegenerator](https://docs.apiverve.com/ref/qrcodegenerator?utm_source=npm&utm_medium=readme).
You can find parameters, example responses, and status codes documented here.

### Setup

```javascript
const qrcodegeneratorAPI = require('@apiverve/qrcodegenerator');
const api = new qrcodegeneratorAPI({
    api_key: '[YOUR_API_KEY]'
});
```

---

## Usage

---

### Perform Request

Using the API is simple. All you have to do is make a request. The API will return a response with the data you requested.

```javascript
var query = {
  "value": "https://apiverve.com",
  "type": "url",
  "format": "png",
  "margin": "0"
};

api.execute(query, function (error, data) {
    if (error) {
        return console.error(error);
    } else {
        console.log(data);
    }
});
```

---

### Using Promises

You can also use promises to make requests. The API returns a promise that you can use to handle the response.

```javascript
var query = {
  "value": "https://apiverve.com",
  "type": "url",
  "format": "png",
  "margin": "0"
};

api.execute(query)
    .then(data => {
        console.log(data);
    })
    .catch(error => {
        console.error(error);
    });
```

---

### Using Async/Await

You can also use async/await to make requests. The API returns a promise that you can use to handle the response.

```javascript
async function makeRequest() {
    var query = {
  "value": "https://apiverve.com",
  "type": "url",
  "format": "png",
  "margin": "0"
};

    try {
        const data = await api.execute(query);
        console.log(data);
    } catch (error) {
        console.error(error);
    }
}
```

---

## Example Response

```json
{
  "status": "ok",
  "error": null,
  "data": {
    "id": "31d88a02-4dd3-4c97-9725-995551c31512",
    "format": "png",
    "type": "url",
    "correction": "M",
    "size": 5,
    "margin": 0,
    "expires": 1740259974243,
    "downloadURL": "https://storage.googleapis.com/apiverve.appspot.com/qrcodegenerator/7277a6dc-6b1f-4af4-820a-755eb1c1a24d.png?GoogleAccessId=635500398038-compute%40developer.gserviceaccount.com&Expires=1740259974&Signature=kk80QPNRGLhIXm6%2BmKozIhoq2ZpehDsiT7%2FyHp4JxvSj5f6EdktMxJrZipxZiz%2BgQNKdr%2F30jaBk8DWT2Q6rPNdQjbJlh1ibu6FnghwMTwtJwU7Y0o2xBXm1b6p9TdqcpMK3ueJviYJ7mnsxdIud6aIQb1t%2FnaUQM%2Fu4dIB8%2By5xXIzldiKKu7zXttyQYb5qHIb89gadMfjd8SVVJVNXm97eroa82C9QL%2Bspuz9grMvyq7x7PfpXW3pZo89x9QKYQu1hT%2Bv1exID7CKv7keZcQ5v7IdSCPGKZyweGsVRQHJ4gwWHJcA2%2FGXrMUKVzXjkUsYBhEArblnh3FVI2%2FP3FQ%3D%3D"
  }
}
```

---

## Customer Support

Need any assistance? [Get in touch with Customer Support](https://apiverve.com/contact?utm_source=npm&utm_medium=readme).

---

## Updates

Stay up to date by following [@apiverveHQ](https://twitter.com/apiverveHQ) on Twitter.

---

## Legal

All usage of the APIVerve website, API, and services is subject to the [APIVerve Terms of Service](https://apiverve.com/terms?utm_source=npm&utm_medium=readme), [Privacy Policy](https://apiverve.com/privacy?utm_source=npm&utm_medium=readme), and [Refund Policy](https://apiverve.com/refund?utm_source=npm&utm_medium=readme).

---

## License
Licensed under the The MIT License (MIT)

Copyright (&copy;) 2026 APIVerve, and EvlarSoft LLC

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
