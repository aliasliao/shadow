import * as React from 'react'
import { useEffect, useRef, useState } from 'react'
import { noop, http } from './utils'
import cx from 'classnames'
import { cardHeader, s } from './styles'
import { Loading } from './components'

export default function App() {
  return (
    <div className={s.vGap12}>
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
  const [loading, setLoading] = useState(false)

  const getStats = () => http.noop(setLoading(true)).get('/getStats')
    .then(data => { setData(data); setError(null) })
    .catch(error => setError(error))
    .finally(() => setLoading(false))
  useEffect(() => { getStats().then(noop) }, [])

  const stopApp = () => http.noop(setLoading(true)).post('/stopApp')
    .catch(error => setError(error))
    .finally(() => setLoading(false))

  return (
    <div className={cx(s.light, s['vGap-1'])}>
      <div className={cx(
        s.bd, s.bgSecondary, s.fontBolder, s.light, s.p4, s.pl12,
        s.dFlex, s.justifyBetween, s.alignCenter,
        cardHeader,
      )}>
        <div>Stats</div>
        {loading && <Loading />}
      </div>
      {error && (
        <div className={cx(s.bd, s.bgDanger, s.p8, s.pl12)}>{error.message}</div>
      )}
      {data && (
        Object.keys(data).map(key => (
          <div key={key} className={cx(s.bd, s.dFlex, s.flexRow)}>
            <div className={cx(s.p4, s.dFlex, s.alignItemsCenter)} style={{ width: '35%' }}>{key}</div>
            <div className={cx(s.bdLeft, s.p4, s.dFlex, s.alignItemsCenter)}>{data[key]}</div>
          </div>
        ))
      )}
      <div className={cx(s.dFlex, s.justifyCenter, s.hGap8, s.bd, s.p8)}>
        <button onClick={getStats}>Get Stats</button>
        <button onClick={stopApp}>Stop App</button>
      </div>
    </div>
  )
}

function StartApp() {
  const [error, setError] = useState<Error | null>(null)
  const [loading, setLoading] = useState(false)
  const startApp = (values: {
    url: string,
    cache: boolean,
    loglevel: string,
  }) => http.noop(setLoading(true)).post('/startApp', values)
    .catch(error => setError(error))
    .finally(() => setLoading(false))

  const formElement = useRef<HTMLFormElement>(null)

  const handleSubmit = () => {
    // @ts-ignore
    const formValues = Object.fromEntries([...new FormData(formElement.current).entries()])
    startApp({
      ...formValues as any,
      cache: formValues.cache === 'on',
    }).then(noop)
  }

  return (
    <div className={cx(s.light, s['vGap-1'])}>
      <form ref={formElement} className={s['vGap-1']}>
        <div className={cx(
          s.bd, s.bgSecondary, s.fontBolder, s.light, s.p4, s.pl12,
          s.dFlex, s.justifyBetween, s.alignCenter,
          cardHeader,
        )}>
          <div>Start App</div>
          {loading && <Loading />}
        </div>
        {error && (
          <div className={cx(s.bd, s.bgDanger, s.p8, s.pl12)}>{error.message}</div>
        )}
        <div className={cx(s.bd, s.dFlex, s.flexRow)}>
          <div className={cx(s.p4, s.dFlex, s.alignItemsCenter)} style={{ width: '35%' }}>
            <label htmlFor="url">url</label>
          </div>
          <div className={cx(s.bdLeft, s.p4, s.dFlex, s.alignItemsCenter, s.grow)}>
            <input type="text" name="url" className={s.w100} />
          </div>
        </div>
        <div className={cx(s.bd, s.dFlex, s.flexRow)}>
          <div className={cx(s.p4, s.dFlex, s.alignItemsCenter)} style={{ width: '35%' }}>
            <label htmlFor="cache">cache</label>
          </div>
          <div className={cx(s.bdLeft, s.p4, s.dFlex, s.alignItemsCenter)}>
            <input type="checkbox" defaultChecked name="cache" />
          </div>
        </div>
        <div className={cx(s.bd, s.dFlex, s.flexRow)}>
          <div className={cx(s.p4, s.dFlex, s.alignItemsCenter)} style={{ width: '35%' }}>
            <label htmlFor="loglevel" defaultValue="warning">loglevel</label>
          </div>
          <div className={cx(s.bdLeft, s.p4, s.dFlex, s.alignItemsCenter)}>
            <select name="loglevel" defaultValue="warning">
              {['debug', 'info', 'warning', 'error', 'none'].map(level => (
                <option key={level} value={level}>{level}</option>
              ))}
            </select>
          </div>
        </div>
      </form>
      <div className={cx(s.dFlex, s.justifyCenter, s.bd, s.p8)}>
        <button onClick={handleSubmit}>Submit</button>
      </div>
    </div>
  )
}
