> Microkubes Todo

# Todo frontend

## Build Setup

#### Install dependencies
``` 
npm install
```

#### Serve with hot reload at localhost:8080
```
npm run dev
```

#### Build for production with minification
```
npm run build
```

#### Build for production and view the bundle analyzer report
```
npm run build --report
```

#### Run unit tests
```
npm run unit
```

#### Run all tests
```
npm test
```

#### Enable CORS
Due to communicating on different hosts, Cross-origin resource sharing (CORS) should be enabled on Kong and on ```jwt-issuer``` with this command:
```
curl -X POST http://localhost:8001/apis/{service_id}/plugins \
    --data "name=cors" \
    --data "config.origins=*" \  
    --data "config.methods=GET, POST, DELETE"\
    --data "config.max_age=3600"
```

#### Vue.js
Vue.js is an open-source JavaScript framework for building user interfaces. Vue can also function as a web application framework capable of powering advanced single-page applications. It features an incrementally adoptable architecture.
##### Webpack
Webpack is an open-source JavaScript module bundler. Its main purpose is to bundle JavaScript files for usage in a browser, yet it is also capable of transforming, bundling, or packaging just about any resource or asset. Webpack takes modules with dependencies and generates static assets representing those modules.
##### vue-loader
Webpack loader for Vue.js components.

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).