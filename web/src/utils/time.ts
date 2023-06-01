export function TimeFormatter(date: number): string {
  const d = new Date(date * 1000).toJSON()

  const newDate = new Date(+new Date(d) + 8 * 3600 * 1000)
    .toISOString()
    .replace(/T/g, ' ')
    .replace(/\.[\d]{3}Z/, '')
  return newDate
}
