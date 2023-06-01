export function SizeFormatterFromBytes(size: number): string {
  const bytesToMega = size / 1024 ** 2
  const bytesToGiga = size / 1024 ** 3
  return bytesToMega < 1024
    ? String(bytesToMega.toFixed(2) + ' MB')
    : String(bytesToGiga.toFixed(2) + ' GB')
}
