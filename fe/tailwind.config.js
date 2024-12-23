/** @type {import('tailwindcss').Config} */
export default {
  /// ISUUE : https://github.com/tailwindlabs/tailwindcss/discussions/6019#discussion-3673982
  /// SOLVED : https://github.com/tailwindlabs/tailwindcss/discussions/6019#discussioncomment-1609444
  content: ['./index.html', './src/**/*.{vue,js,ts,jsx,tsx}'],
  theme: {
    extend: {},
  },
  plugins: [],
}
