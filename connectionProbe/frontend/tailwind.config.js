/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './app/**/*.{js,ts,jsx,tsx}',
    './components/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      keyframes: {
        dots: {
          '0%': { width: '0px' },
          '33%': { width: '6px' },
          '66%': { width: '12px' },
          '100%': { width: '18px' }
        }
      },
      animation: {
        dots: 'dots 1s steps(4) infinite'
      }
    },
  },
  plugins: [],
}
