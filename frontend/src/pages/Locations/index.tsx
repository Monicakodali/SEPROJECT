import * as React from 'react';
import Map from './Map'
import LocationsList from './LocationsList'
import FilterDrawer from './FilterDrawer';
import { Grid } from '@mui/material';
import { styled } from '@mui/material/styles';
import dining from './dining.json'
import { LatLngTuple } from 'leaflet';

const locations = dining.features.map((f) => {
  return {
    id: parseInt(f.properties.ID, 10),
    name: f.properties.Name,
    building: f.properties.BLDG,
    room: f.properties.ROOM,
    coordinates: [f.geometry.coordinates[1], f.geometry.coordinates[0]] as LatLngTuple
  }
})

const FILTER_DRAWER_WIDTH = 225

const LocationListContainer = styled('div')(({ theme }) => ({
  marginLeft: FILTER_DRAWER_WIDTH,
  padding: theme.spacing(2)
}));


export default function Locations() {
  
  return (
    <Grid container sx={{flexGrow: 1, maxHeight: '100%'}}>
      <Grid item xs={12} sm={7} sx={{position: 'relative', overflow: 'auto'}}>
        <FilterDrawer drawerWidth={FILTER_DRAWER_WIDTH}/>
        <LocationListContainer>
          <LocationsList locations={locations} />
        </LocationListContainer>
      </Grid>
      <Grid item xs={12} sm={5} sx={{height: '100%'}}>
        <Map locations={locations} />  
      </Grid>
    </Grid>
  );
};