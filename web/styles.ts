// @ts-ignore
import { css } from 'emotion'

const low = str => (str.charAt(0).toLowerCase() + str.substring(1))
const up = str => (str.charAt(0).toUpperCase() + str.substring(1))
//

export const s: {
  [key: string]: string,
} = {
  ...Object.fromEntries(
    ['p', 'm'].map(
      type => [-4,-2,-1,1,2,4,8,12,16].map(size => {
        const t = { p : 'padding', m: 'margin' }[type]
        return Object.entries({
          [`${type}${size}`]: css`${t}: ${size}px`,
          [`${type}x${size}`]: css`${t}-left: ${size}px; ${t}-right: ${size}px`,
          [`${type}y${size}`]: css`${t}-top: ${size}px; ${t}-bottom: ${size}px`,
          [`${type}t${size}`]: css`${t}-top: ${size}px`,
          [`${type}b${size}`]: css`${t}-bottom: ${size}px`,
          [`${type}l${size}`]: css`${t}-left: ${size}px`,
          [`${type}r${size}`]: css`${t}-right: ${size}px`,
        })
      })
    ).flat(2)
  ),
  ...Object.fromEntries(
    [2,4,8,12,16].map(size => [`bdRadius${size}`, css`border-radius: ${size}px`])
  ),
  ...Object.fromEntries(
    ['', 'bd', 'bg'].map(type => {
      const t = { '': '', bd: 'border-', bg: 'background-' }[type]
      return Object.entries({
        [`${type}Primary`]: css`${t}color: #007bff`,
        [`${type}Secondary`]: css`${t}color: #6c757d`,
        [`${type}Success`]: css`${t}color: #28a745`,
        [`${type}Danger`]: css`${t}color: #dc3545`,
        [`${type}Warning`]: css`${t}color: #ffc107`,
        [`${type}Info`]: css`${t}color: #17a2b8`,
        [`${type}Light`]: css`${t}color: #f8f9fa`,
        [`${type}Dark`]: css`${t}color: #343a40`,
      }).map(([k, v]) => [low(k), v])
    }).flat()
  ),
  ...Object.fromEntries(
    [10,11,12,13,14,15,16,18,20].map(
      size => [`font${size}`, `${size}px`]
    )
  ),
  ...Object.fromEntries(
    ['lighter','normal','bold','bolder'].map(
      weight => [`font${up(weight)}`, css`font-weight: ${weight}`],
    )
  ),
  ...{
    dFlex: css`display: flex`,
    flexColumn: css`flex-direction: column`,
    flexRow: css`flex-direction: row`,
    grow: css`flex-grow: 1`,
    shrink: css`flex-shrink: 1`,
    w100: css`width: 100%`,
    // border
    bd: css`border: 1px solid`,
    bdTop: css`border-top: 1px solid`,
    bdBottom: css`border-bottom: 1px`,
    bdLeft: css`border-left: 1px solid`,
    bdRight: css`border-right: 1px solid`,
  },
}

console.log(s)
