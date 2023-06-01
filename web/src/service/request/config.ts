let BASE_URL = ''
const TIME_OUT = 10000

if (process.env.NODE_ENV === 'development') {
  BASE_URL = 'http://192.168.251.176:9527/v1'
} else if (process.env.NODE_ENV === 'production') {
  BASE_URL = '/v1'
}

export { BASE_URL, TIME_OUT }
