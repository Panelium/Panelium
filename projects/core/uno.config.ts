import {defineConfig} from "unocss";
import presetWind3 from "npm:@unocss/preset-wind3";

export default defineConfig({
    presets: [
        presetWind3(),
    ],
    cli: {
        entry: {
            patterns: ["./internal/dashboard/views/**/*.{html,templ,go}"],
            outFile: "./internal/dashboard/static/css/vendor/uno.min.css",
        },
    },
});
