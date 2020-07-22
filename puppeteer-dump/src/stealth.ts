import puppeteer from 'puppeteer-extra';
import StealthPlugin from 'puppeteer-extra-plugin-stealth';
import * as fs from 'fs';
import * as jsbeautifier from 'js-beautify';
import * as debug from 'debug';

const DIST_FILE = 'dist/stealth.js';
let CDP_MSG: string[] = new Array();

// enable debug
debug.enable('puppeteer:protocol');

// https://stackoverflow.com/questions/32719923/redirecting-stdout-to-file-nodejs/60027895#60027895
const stderrWrite = process.stderr.write;
process.stderr.write = (args: string) => { // On stderr write

  CDP_MSG.push(args);

  return stderrWrite.apply(process.stderr, [args]);
};

puppeteer
  .use(StealthPlugin())
  .launch({
    headless: true,
  })
  .then(async browser => {
    const page = await browser.newPage();
    await page.goto('about:blank');
    await browser.close();

    const regex = /^[^\{\}\n]+(\{[^\n]+\})[^\{\}]+$/m;

    fs.writeFileSync(DIST_FILE, jsbeautifier.js_beautify(CDP_MSG
      .filter(msg => regex.test(msg))
      .map(msg => {
        if (regex.test(msg)) {
          let data = JSON.parse(msg.replace(regex, '$1'));
          return data &&
            data.method == 'Page.addScriptToEvaluateOnNewDocument' &&
            data.params.worldName != '__puppeteer_utility_world__' &&
            data.params.source ? data.params.source + ';' : '';
        }
      })
      .join('\n\n')));
  });






