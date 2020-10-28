import { notificationInstance, s, notificationContainer } from './styles'

export const HOST = 'http://router.asus.com:3000'

function factory(
  method: 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE'
): (
  (url: string, body?: any) => Promise<any>
) {
  return (url, body = '') => fetch(`${HOST}${url}`, {
    method,
    body: method === 'GET' ? null : JSON.stringify(body),
    credentials: "include",
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
      'Accept': 'application/json;charset=utf-8',
    },
  }).then((res) => {
    if (!res.ok) {
      return res.text().then(reason => Promise.reject(new Error(reason)))
    }
    if (res.headers.get('Content-Type').includes('json')) {
      return res.json()
    }
    return res.text()
  }).then((data: any) => {
    notify({ message: 'Operation Succeeded 😁' })
    return data
  }, (error: Error) => {
    notify({ message: 'Operation Failed 😞', type: 'danger' })
    throw error
  })
}

export const http = {
  noop: (_: any) => http,
  get: factory('GET'),
  post: factory('POST'),
  put: factory('PUT'),
  patch: factory('PATCH'),
  delete: factory('DELETE'),
}
export const noop = () => {}
const up = str => (str.charAt(0).toUpperCase() + str.substring(1))

export const initializeNotification = () => {
  const container = document.createElement('div')
  container.id = 'notificationContainer'
  container.classList.add(notificationContainer)
  document.body.appendChild(container)
}

export const notify = ({ message, type = 'info' }: {
  message: string,
  type?: 'primary' | 'secondary' | 'success' | 'danger' | 'warning' | 'info' | 'dark'
}) => {
  const container = document.getElementById('notificationContainer')

  const span = document.createElement('span')
  span.innerText = message
  span.classList.add(notificationInstance, s.light, s[`bg${up(type)}`])
  container.appendChild(span)

  window.setTimeout(() => {
    container.removeChild(span)
  }, 3000)
}
