# bypass

> anti-detect with rod

## Overview

This is a side project of [`rod`](https://github.com/go-rod/rod).

It will dump the raw js code which is injected to the puppeteer by [puppeteer-extra-plugin-stealth](https://github.com/berstend/puppeteer-extra/tree/master/packages/puppeteer-extra-plugin-stealth), then we can use them in any [cdp](https://chromedevtools.github.io/devtools-protocol/) implementation.

## Pure JS

You can use the js code anywhere, just fetch it from: 

`https://cdn.jsdelivr.net/gh/go-rod/bypass@latest/puppeteer-dump/dist/stealth.js`


## Reference

- [puppeteer-typescript-boilerplate](https://github.com/sosmii/puppeteer-typescript-boilerplate)
- [puppeteer-extra-plugin-stealth](https://github.com/berstend/puppeteer-extra/tree/master/packages/puppeteer-extra-plugin-stealth)
- [puppeteer debugging tips](https://github.com/puppeteer/puppeteer#debugging-tips)