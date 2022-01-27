import * as React from 'react';
import Map from './Map'
import LocationsList from './LocationsList'
import Filters from './Filters';
import { styled } from '@mui/material/styles';
import { LatLngTuple } from 'leaflet';
import { InsetDrawer } from '../../components';

// the hardcoded dining locations JSON file is temporary
// @TODO: move these restaurant locations to the DB
import dining from './dining.json'

const locations = dining.features.map((f) => {
  return {
    id: parseInt(f.properties.ID, 10),
    name: f.properties.Name,
    building: f.properties.BLDG,
    room: f.properties.ROOM,
    coordinates: [f.geometry.coordinates[1], f.geometry.coordinates[0]] as LatLngTuple
  }
})

// rework to make map a percentage of screen
const FILTER_DRAWER_WIDTH = 225
const MAP_DRAWER_WIDTH = '40%'

const LocationListContainer = styled('div')(({ theme }) => ({
  marginLeft: FILTER_DRAWER_WIDTH,
  marginRight: MAP_DRAWER_WIDTH,
  padding: theme.spacing(2),
}));


export default function Locations() {
  
  return (
    <div>
      <InsetDrawer anchor="left" width={FILTER_DRAWER_WIDTH}>
        <Filters />
      </InsetDrawer>
      <LocationListContainer>
        <LocationsList locations={locations} />
      </LocationListContainer>
      <InsetDrawer anchor="right" width={MAP_DRAWER_WIDTH}>
        <Map locations={locations} />  
      </InsetDrawer>
    </div>
  );
};