/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./Frontend/**/*.html",
    "./src/**/*.{html,js,jsx,ts,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        cream: "#F4F1E8",
        redbrown: "#986141",
      },
    },
  },
  plugins: [],
}
