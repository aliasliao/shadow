import * as React from 'react'
import { useEffect, useState } from 'react'
import { noop, http } from './utils'
import cx from 'classnames'
import { s } from './styles'

export default function App() {
  return (
    <>
      <Stats />
      <StartApp />
    </>
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

  return (
    <div className={cx(s.w100, s.light)}>
      <div className={cx(s.bgSecondary, s.bd, s.bdDark, s.fontBolder, s.light, s.p4, s.pl12)}>Stats</div>
      {data === null && error === null && (
        <div className={cx(s.p8, s.pl12, s.bgInfo)}>Loading Status...</div>
      )}
      {error && (
        <div className={cx(s.p8, s.bgDanger)}>{error.message}</div>
      )}
      {data && (
        Object.keys(data).map(key => (
          <div key={key} className={cx(s.bd, s.bdDark, s.dFlex, s.flexRow)}>
            <div className={cx(s.p4)} style={{ width: '35%' }}>{key}</div>
            <div className={cx(s.bdLeft, s.bdDark, s.p4)}>{data[key]}</div>
          </div>
        ))
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
      {data === null && error === null && (
        <>Starting App...</>
      )}
      <button type="submit">Submit</button>
    </form>
  )
}
