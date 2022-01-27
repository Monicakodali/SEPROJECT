import { LatLngTuple } from 'leaflet';

export type Location = {
  name: string
  id: number,
  building: string,
  room: string,
  coordinates: LatLngTuple
}