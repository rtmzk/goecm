import type { UploadRequestOptions } from 'element-plus'
import { postFileUpload } from '@/service/image/image'
import SparkMD5 from 'spark-md5'
import { ref, computed } from 'vue'

export default function upload(option: UploadRequestOptions) {
  const fileChunkList = []
  const chunkSize = 5 * 1024 * 1024
  const chunkFormData = ref<any>()
  const fileHash = ref()
  const percentage = computed(() => {
    if (!chunkFormData.value?.length) return 0
    let uploaded = chunkFormData.value.filter(
      (item: any) => item.percentage
    ).length
    return Number(((uploaded / chunkFormData.value.length) * 100).toFixed(2))
  })

  let cur = 0
  while (cur < option.file.size) {
    fileChunkList.push(option.file.slice(cur, cur + chunkSize))
    cur += chunkSize
  }

  fileHash.value = createMD5(fileChunkList)
  let chunkList = fileChunkList.map((file, index) => {
    return {
      file: file,
      chunkHash: fileHash.value + '-' + index,
      fileHash: fileHash.value
    }
  })

  chunkFormData.value = chunkList.map((chunk, index) => {
    let formData = new FormData()
    formData.append('file', chunk.file)
    formData.append('chunk', String(index + 1))
    formData.append('chunks', String(fileChunkList.length))
    formData.append('fileName', option.file.name)

    return {
      formData: formData,
      percentage: 0
    }
  })

  if (chunkFormData.value.length > 0) {
    for (let i = 0; i < chunkFormData.value.length; i++) {
      postFileUpload(chunkFormData.value[i].formData).then((res: any) => {
        if (res.status === 'uploading') {
          uploadProgress(chunkFormData.value[i])
        }
      })
    }
  }

  // chunkFormData.value.map((data) => {
  //   return new Promise((resolve, reject) => {
  //     postFileUpload(data.formData)
  //       .then((data) => {
  //         resolve(data)
  //       })
  //       .catch((err) => {
  //         reject(err)
  //       })
  //   })
  // })
}

const uploadProgress = (item: any) => {
  return (e: any) => {
    item.percentage = parseInt(String((e.loaded / e.total) * 100))
  }
}

const createMD5 = (fileChunkList: Array<any>) => {
  return new Promise((resolve, reject) => {
    // const slice = File.prototype.slice
    const chunks = fileChunkList.length
    let currentChunk = 0
    let spark = new SparkMD5.ArrayBuffer()
    let fileReader = new FileReader()

    fileReader.onload = function (e) {
      spark.append(e.target?.result as ArrayBuffer)
      currentChunk++
      if (currentChunk < chunks) {
        loadChunk()
      } else {
        resolve(spark.end)
      }
    }
    fileReader.onerror = function (e) {
      reject(e)
    }
    function loadChunk() {
      fileReader.readAsArrayBuffer(fileChunkList[currentChunk])
    }

    loadChunk()
  })
}
