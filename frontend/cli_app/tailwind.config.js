/** @type {import('tailwindcss').Config} */
export default {
  content: ['./index.html', './src/**/*.{js,ts,jsx,tsx}','./node_modules/flowbite/**/*.js'],
  theme: {
    extend: {
      fontFamily: {
        'roboto':['Roboto', 'sans-serif'],
        'lato':['Lato', 'sans-serif']
      },
    },
  },
    plugins: [
        require('flowbite/plugin')
    ]
}
