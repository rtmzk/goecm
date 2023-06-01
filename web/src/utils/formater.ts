export function TimeFormatter(date: number): string {
  const d = new Date(date * 1000).toJSON()

  const newDate = new Date(+new Date(d) + 8 * 3600 * 1000)
    .toISOString()
    .replace(/T/g, ' ')
    .replace(/\.[\d]{3}Z/, '')
  return newDate
}

export function SizeFormatterFromBytes(size: number): string {
  const bytesToMega = size / 1024 ** 2
  const bytesToGiga = size / 1024 ** 3
  return bytesToMega < 1024
    ? String(bytesToMega.toFixed(2) + ' MB')
    : String(bytesToGiga.toFixed(2) + ' GB')
}

export function StackFormatter(data: string): string {
  const defaultStackName = '-'
  if (data === undefined) {
    return defaultStackName
  }
  return data
}
