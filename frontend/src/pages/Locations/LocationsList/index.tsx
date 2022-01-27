import * as React from 'react';

import LocationCard from './LocationCard'
import { Location } from '../../../types/location'

type LocationsListProps = {
  locations: Location[]
}

export default function LocationsList({locations}: LocationsListProps) {

  const [selected, setSelected] = React.useState<number | null>(null)

  return <>{
    locations.map((d) => <LocationCard selected={selected === d.id} key={d.id} data={d} onClick={() => setSelected(d.id)}/>)
  }
  </>

};