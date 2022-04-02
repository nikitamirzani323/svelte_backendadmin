module.exports = {
  content: ["./src/**/*.{html,js,svelte}"],
  theme: {
    extend: {},
  },
  plugins: [require("daisyui")],
  daisyui: {
    themes: [
      {
        light:{
          ...require("daisyui/src/colors/themes")["[data-theme=light]"],
          primary: "#1a73e8",
        }
      }
    ],
  },
}