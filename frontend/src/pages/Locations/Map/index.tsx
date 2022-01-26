import * as React from 'react';
import { MapContainer, Marker, Popup, TileLayer } from 'react-leaflet'
import L, { LatLngTuple, LocationEvent } from 'leaflet'
import 'leaflet/dist/leaflet.css'
import './index.css'
import { Location } from '../../../types/location'
import { LeafletMouseEvent } from 'leaflet';

type MapProps = {
  locations: Location[]
}

const position: LatLngTuple = [29.6382, -82.3566]

var currentLocationIcon = L.divIcon({
  className: 'loc',
  html: '⭐',
  iconSize: [36, 36],
});

var PoiIcon = L.divIcon({
  className: 'loc',
  html: '📍',
  iconSize: [36, 36],
});


export default function Map({locations}: MapProps) {

  const [userLocation, setUserLocation] = React.useState<LatLngTuple | null>(null)

  return (
      <MapContainer center={position} zoom={14} style={{height: '100%'}} doubleClickZoom={false} whenCreated={(map) => {
        
        map.on('click', (e: LeafletMouseEvent) => {
          console.log(e)
          //setUserLocation([e.latlng.lat, e.latlng.lng])
        })

        map.on('locationfound', (e: LocationEvent) => {
          console.log(e)
          setUserLocation([e.latlng.lat, e.latlng.lng])
        })
        map.locate()
      }}>
      <TileLayer
        attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
        url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
      />
      {/* <Marker position={position}>
        <Popup>
          A pretty CSS3 popup. <br /> Easily customizable.
        </Popup>
      </Marker> */}
      {userLocation && <Marker position={userLocation} icon={currentLocationIcon}/>}
      {
        locations.map(({id, coordinates, name}) => {
          return (<Marker key={id} position={coordinates} icon={PoiIcon}>
            <Popup>
              {name}
            </Popup>
          </Marker>)
        })
      }
  </MapContainer>
  );
};