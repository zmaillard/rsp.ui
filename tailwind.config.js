/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'media',
  content: ["./layouts/**/*.{html,js}", "./node_modules/flowbite/**/*.js"],
  theme: {
    extend: {},
  },
  plugins: [
      require('flowbite/plugin'),
      require('@tailwindcss/typography')
  ],
}

