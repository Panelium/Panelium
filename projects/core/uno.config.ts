import {defineConfig} from "unocss";

export default defineConfig({
    cli: {
        entry: {
            patterns: ["./internal/dashboard/views/**/*.{html,templ,go}"],
            outFile: "./internal/dashboard/static/css/vendor/uno.min.css",
        },
    },
});
