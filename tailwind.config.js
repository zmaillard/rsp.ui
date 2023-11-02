/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'media',
  content: ["./layouts/**/*.{html,js}"],
  theme: {
    extend: {},
  },
  plugins: [require('flowbite/plugin')],
}

