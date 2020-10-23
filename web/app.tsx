import * as React from 'react'
import { useEffect, useState } from 'react'
import { noop, http } from './utils'

export default function App() {
  return (
    <div>
      <Stats />
      <StartApp />
    </div>
  )
}

function Stats() {
  const [data, setData] = useState<{
    NumGoroutine: number,
    NumGC: number,
    Alloc: string,
    TotalAlloc: string,
    Sys: string,
    Mallocs: string,
    Frees: string,
    LiveObjects: string,
    Uptime: number
  } | null>(null)
  const [error, setError] = useState<Error | null>(null)
  const getStats = () => http.get('/getStats')
    .then(data => setData(data))
    .catch(error => setError(error))
  useEffect(() => { getStats().then(noop) }, [])

  if (data === null && error === null) {
    return <>Loading status...</>
  }
  return (
    <div>
      {error && (
        <div>{error.message}</div>
      )}
      {data && (
        <div>
          {Object.keys(data).map(key => (
            <div key={key} className="d-flex">
              <div>{key}</div>
              <div>{data[key]}</div>
            </div>
          ))}
        </div>
      )}
      <button onClick={() => getStats()}>Get Stats</button>
    </div>
  )
}

function StartApp() {
  const [data, setData] = useState<{} | null>(null)
  const [error, setError] = useState<Error | null>(null)
  const startApp = (values: {
    url: string,
    cache: boolean,
    loglevel: string,
  }) => http.post('/startApp', {})
    .then(data => setData(data))
    .catch(error => setError(error))

  if (data === null && error === null) {
    return <>Starting App...</>
  }
  return (
    // @ts-ignore
    <form onSubmit={e => startApp(e.target)}>
      <div>
        <label htmlFor="url">url</label>
        <input type="text" id="url" />
      </div>
      <div>
        <label htmlFor="cache">cache</label>
        <input type="checkbox" id="cache" />
      </div>
      <div>
        <label htmlFor="loglevel">loglevel</label>
        <select id="loglevel">
          {['warning', 'error', 'debug'].map(level => (
            <option key={level} value={level}>{level}</option>
          ))}
        </select>
      </div>
      <button type="submit">Submit</button>
    </form>
  )
}
