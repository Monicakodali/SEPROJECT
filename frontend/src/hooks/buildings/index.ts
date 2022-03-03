import React from "react"
import buildings from './boundaries.json'

const useBuilding = (e: Establishment | null) => {

  return React.useMemo(() => {
    if(!e) {
      return null
    }
    const building = buildings.features.find(f => f.properties.PropCID === e.building)?.properties
    return building ?? null
  }, [e])

}

export default useBuilding