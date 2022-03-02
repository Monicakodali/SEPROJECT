import React from "react"
import hours from './spring2022.json'
import buildings from './boundaries.json'

const useDiningHours = (e: Establishment | null) => {

  return React.useMemo(() => {
    if(!e) {
      return null
    }
    const building = buildings.features.find(f => f.properties.PropCID === e.building)?.properties
    const hour = hours.data.find(h => {
      return h.name.includes(e.name) && (building?.PropName || '').split(' ').some((word: string) => h.fullName.includes(word))
    })
    if(!hour || !Array.isArray(hour.hours)) {
      return null
    }
    return hour?.hours ?? null
  }, [e])

}

export default useDiningHours