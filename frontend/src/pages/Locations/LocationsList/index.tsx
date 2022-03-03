import * as React from 'react';

import LocationCard from './LocationCard'

type LocationsListProps = {
  locations: Establishment[]
}

export default function LocationsList({locations}: LocationsListProps) {

  const [selected, setSelected] = React.useState<string | null>(null)

  return <>{
    locations.map((d) => <LocationCard selected={selected === d.id} key={d.id} data={d} onClick={() => setSelected(d.id)}/>)
  }
  </>

};