import { Skeleton, Typography } from '@mui/material';
import * as React from 'react';
import { styled } from '@mui/material/styles';
import LocationCard from './LocationCard'
import { Box } from '@mui/system';

type LocationsListProps = {
  locations: Establishment[],
  loading: boolean
}

const SkeletonContainer = styled('div')(({ theme }) => {

  return {
    borderWidth: 1,
    borderColor: theme.palette.grey[300],
    borderStyle: 'solid',
    padding: theme.spacing(3),
    '&:not(:last-child)': {
      marginBottom: theme.spacing(2)
    },
  }
})


export default function LocationsList({locations, loading}: LocationsListProps) {

  const [selected, setSelected] = React.useState<string | null>(null)

  if(!loading && locations.length === 0) {
    return <Box sx={{py: 2}}>
      <Typography sx={{color: 'grey.600'}}>There are no locations found with these criteria.</Typography>
    </Box>
  }

  if(loading) {
    return <>{
    [1,2,3,4,5,6].map(a => <SkeletonContainer key={a}>
        <Skeleton sx={{ my: 1 }}/>
        <Skeleton sx={{ my: 1 }}/>
        <Skeleton sx={{ my: 1 }}/>
    </SkeletonContainer>)}
    </>
  }

  return <>{
    locations.map((d) => <LocationCard selected={selected === d.id} key={d.id} data={d} onClick={() => setSelected(d.id)}/>)
  }
  </>

};