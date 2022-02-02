import * as React from 'react';
import Map from './Map'
import LocationsList from './LocationsList'
import Filters from './Filters';
import { styled } from '@mui/material/styles';
import { InsetDrawer } from '../../components';
import axios from 'axios';


const FILTER_DRAWER_WIDTH = 225
const MAP_DRAWER_WIDTH = '40%'

const LocationListContainer = styled('div')(({ theme }) => ({
  marginLeft: FILTER_DRAWER_WIDTH,
  marginRight: MAP_DRAWER_WIDTH,
  padding: theme.spacing(2),
}));


export default function Locations() {
  
  const [locations, setLocations] = React.useState([])

  React.useEffect(() => {
    axios.get('/api/establishments').then(res => {
      setLocations(res.data)
    })
  }, [])

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