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
  })
}

export const http = {
  noop: (value: any) => http,
  get: factory('GET'),
  post: factory('POST'),
  put: factory('PUT'),
  patch: factory('PATCH'),
  delete: factory('DELETE'),
}
export const noop = () => {}
