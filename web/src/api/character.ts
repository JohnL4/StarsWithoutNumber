export interface Character {
  ST: number
  DX: number
  CO: number
  IQ: number
  WI: number
  CH: number
}

export function emptyCharacter(): Character {
  return { ST: 0, DX: 0, CO: 0, IQ: 0, WI: 0, CH: 0 }
}
