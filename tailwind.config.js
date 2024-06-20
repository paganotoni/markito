module.exports = {
  content: [
    'internal/*.html',
    'internal/**/*.html',
    "public/*.js",
  ],

  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],

  theme: {
    extend: {
      maxHeight: {
        'inherit': 'inherit',
      }
    },
  },
}