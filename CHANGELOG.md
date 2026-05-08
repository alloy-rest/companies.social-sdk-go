# Changelog

## 0.2.0 (2026-05-08)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/alloy-rest/companies.social-sdk-go/compare/v0.1.0...v0.2.0)

### Features

* **go:** add default http client with timeout ([7e7ffcd](https://github.com/alloy-rest/companies.social-sdk-go/commit/7e7ffcd9e54f60f20f68d11d2b534dc0d4c0e286))
* **internal:** support comma format in multipart form encoding ([6b73298](https://github.com/alloy-rest/companies.social-sdk-go/commit/6b732986a663116e49938c789cdf691c36197326))
* support setting headers via env ([8f8fcf3](https://github.com/alloy-rest/companies.social-sdk-go/commit/8f8fcf3d1589d20e22e1a93ea18622bebe5853dd))


### Bug Fixes

* **go:** avoid panic when http.DefaultTransport is wrapped ([3d44d4c](https://github.com/alloy-rest/companies.social-sdk-go/commit/3d44d4c737cb658dd09140f83843783fe091aaba))
* prevent duplicate ? in query params ([5b84995](https://github.com/alloy-rest/companies.social-sdk-go/commit/5b8499525d4c27c1770d30ada3e27d2b4a22ae23))


### Chores

* avoid embedding reflect.Type for dead code elimination ([467fb46](https://github.com/alloy-rest/companies.social-sdk-go/commit/467fb46c5701cf8a2f1798e7f779316cebf6fdc8))
* **ci:** skip lint on metadata-only changes ([4e89e8e](https://github.com/alloy-rest/companies.social-sdk-go/commit/4e89e8e8d13797daeb4c18ae13d744694aa4cd2e))
* **ci:** support opting out of skipping builds on metadata-only commits ([07ead66](https://github.com/alloy-rest/companies.social-sdk-go/commit/07ead66fe67615de6d36f5e4fa748e1297a32cd5))
* **client:** fix multipart serialisation of Default() fields ([8cd3f24](https://github.com/alloy-rest/companies.social-sdk-go/commit/8cd3f24c1345f979b002b1da3be65d0230b5be1c))
* **internal:** more robust bootstrap script ([0c99dea](https://github.com/alloy-rest/companies.social-sdk-go/commit/0c99dea82d7441085962cf8f7d8f6cf073f1bd79))
* **internal:** support default value struct tag ([169e068](https://github.com/alloy-rest/companies.social-sdk-go/commit/169e0685bbe6cf43712b024bfe99e0760092d21f))
* **internal:** tweak CI branches ([b8c9f45](https://github.com/alloy-rest/companies.social-sdk-go/commit/b8c9f4506b70a5689c859a4d26c8e54d82999ee2))
* **internal:** update gitignore ([8b2e623](https://github.com/alloy-rest/companies.social-sdk-go/commit/8b2e62384d7921f6f6b0c8d3fead18e035c5066c))
* redact api-key headers in debug logs ([aefadff](https://github.com/alloy-rest/companies.social-sdk-go/commit/aefadff714955b436f90479027203bfa35937ab6))
* remove unnecessary error check for url parsing ([0f19e07](https://github.com/alloy-rest/companies.social-sdk-go/commit/0f19e07e619a2f809d71a57de08f50af00cf0259))
* update docs for api:"required" ([308d47f](https://github.com/alloy-rest/companies.social-sdk-go/commit/308d47f3cd8c83b9059bd3710950ae48edce84b0))

## 0.1.0 (2026-03-11)

Full Changelog: [v0.0.2...v0.1.0](https://github.com/alloy-rest/companies.social-sdk-go/compare/v0.0.2...v0.1.0)

### Features

* **api:** api update ([8879be0](https://github.com/alloy-rest/companies.social-sdk-go/commit/8879be0993eb6234f24a1593ed43eb4d12cdee09))
* **api:** manual updates ([b2ef312](https://github.com/alloy-rest/companies.social-sdk-go/commit/b2ef312081469bb4cb53588a440d9d5e098c0a8b))


### Chores

* **internal:** minor cleanup ([263a0ec](https://github.com/alloy-rest/companies.social-sdk-go/commit/263a0ec1523b0c3a15ff52d02c13b0bc1b16ac3a))
* **internal:** use explicit returns ([402d2d5](https://github.com/alloy-rest/companies.social-sdk-go/commit/402d2d5e97f370cfc02640767815dd02451e834d))
* **internal:** use explicit returns in more places ([a790276](https://github.com/alloy-rest/companies.social-sdk-go/commit/a7902769a2c999ba03c058ab6f7dbe31b9bd60af))

## 0.0.2 (2026-03-08)

Full Changelog: [v0.0.1...v0.0.2](https://github.com/alloy-rest/-companies.social-sdk-go/compare/v0.0.1...v0.0.2)

### Chores

* configure new SDK language ([ae0a6a6](https://github.com/alloy-rest/-companies.social-sdk-go/commit/ae0a6a64cc23a270d35db08aa1f867869f59b350))
* update SDK settings ([2fc8510](https://github.com/alloy-rest/-companies.social-sdk-go/commit/2fc851088651ee8a5c5560a54373a62e89b278d9))
