import * as React from 'react';
import Map from './Map'
import LocationsList from './LocationsList'
import ListHeader from './ListHeader'
import Filters from './Filters';
import { styled } from '@mui/material/styles';
import { InsetDrawer } from '../../components';
import axios from 'axios';
import { useSearchParams } from 'react-router-dom';


const FILTER_DRAWER_WIDTH = 225
const MAP_DRAWER_WIDTH = '40%'

const LocationListContainer = styled('div')(({ theme }) => ({
  marginLeft: FILTER_DRAWER_WIDTH,
  marginRight: MAP_DRAWER_WIDTH,
  padding: theme.spacing(2),
}));


export default function Locations() {
  
  const [locations, setLocations] = React.useState([])
  const [loading, setLoading] = React.useState(true)
  
  let [searchParams] = useSearchParams()
  const query = searchParams.get('query')
  const searchType = searchParams.get('type')

  // currently filtering on the front end
  // need to move search logic to backend
  React.useEffect(() => {
    setLoading(true)
    axios.get('/api/establishments').then(res => {
      console.log(res.data)
      //setLocations(res.data)
      if(query) {
        setLocations(res.data.filter((d: Establishment) => {
          return d?.name?.replace(/[^a-zA-Z0-9]/g, '')?.toLowerCase().includes(query)
        }))
      } else {
        setLocations(res.data)
      }

      setTimeout(() => {
        setLoading(false)
      }, 1000)
    })
  }, [query])

  return (
    <div>
      <InsetDrawer anchor="left" width={FILTER_DRAWER_WIDTH}>
        <Filters />
      </InsetDrawer>
      <LocationListContainer>
        <ListHeader />
        <LocationsList locations={locations} loading={loading} />
      </LocationListContainer>
      <InsetDrawer anchor="right" width={MAP_DRAWER_WIDTH}>
        <Map locations={locations} />  
      </InsetDrawer>
    </div>
  );
};