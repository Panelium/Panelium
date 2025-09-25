import {defineConfig, presetWind3} from 'unocss'

export default defineConfig({
    content: {
        filesystem: ['**/*.{html,templ}'],
    },
    presets: [presetWind3()],
})
