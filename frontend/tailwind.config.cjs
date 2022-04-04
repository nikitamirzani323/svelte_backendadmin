module.exports = {
  content: ["./src/**/*.{html,js,svelte}"],
  theme: {
    extend: {
      animation: {
        'spin-slow-top': 'top 2s linear infinite',
        'spin-slow-bottom': 'bottom 2s linear infinite',
        'spin-slow-right': 'right 2s linear infinite',
        'spin-slow-left': 'left 2s linear infinite',
        'typing-slow': 'typing 5s steps(10) infinite'
      },
      keyframes: {
        top: {
          '0%': { transform:'translatex(-100%)' },
          '50%, 100%': { transform:'translatex(200%)' },
        },
        bottom: {
          '0%': { transform:'translatex(100%)' },
          '50%, 100%': { transform:'translatex(-200%)' },
        },
        right: {
          '0%': { transform:'translatey(-100%)' },
          '50%, 100%': { transform:'translatey(200%)' },
        },
        left: {
          '0%': { transform:'translatey(100%)' },
          '50%, 100%': { transform:'translatey(-200%)' },
        },
        typing: {
          '0%,90%,100%': { width: '0px' },
          '30%,60%': { width: '124px' },
        },
      }
    },
  },
  plugins: [require('tailwind-scrollbar'),require("daisyui"),require("tailwindcss-animation-delay")],
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