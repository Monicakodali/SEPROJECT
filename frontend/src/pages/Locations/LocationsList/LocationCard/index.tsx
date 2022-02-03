import * as React from 'react';
import { Box, CardContent, Typography, CardMedia, ListItemButton } from '@mui/material';
import { styled } from '@mui/material/styles';

import { Location } from '../../../../types/location'
import Rating from './Rating'
import Tags from './Tags'

type LocationsCardProps = {
  data: Location,
  selected?: boolean,
  onClick: () => void
}

const placeholderImageUrl = 'https://thelivingstonpost.com/wp-content/themes/fox/images/placeholder.jpg'

const StyledListItem = styled(ListItemButton)(({ theme }) => {

  return {
    display: 'flex',
    borderWidth: 1,
    borderColor: theme.palette.grey[300],
    borderStyle: 'solid',
    padding: 0,
    '&:not(:last-child)': {
      marginBottom: theme.spacing(2)
    },
    '&:hover, &:focus': {
      boxShadow: theme.shadows[3],
      //backgroundColor: 'initial' //modify hover bg color?
    }
  }
})

const names = ['', 'John', 'Mary', 'Sue', 'Bob', 'Mike']

export default function LocationsCard({data, selected, onClick}: LocationsCardProps) {

  const { name, id } = data

  // @TODO: populate with actual rating
  // randomly generate rating for now
  const rating = React.useMemo(() => Math.floor(Math.random() * 5) + 1, [])

  // @TODO: populate with actual num ratings
  // randomly generate num of ratings for now
  const numRatings = React.useMemo(() => Math.floor(Math.random() * 100) + 10, [])

  // @TODO: populate with actual isOpen
  // randomly generate whether open or not
  const isOpen = React.useMemo(() => Math.random() > 0.5, [])

  // @TODO: populate tags from DB
  const tags = ['Coffee & Tea', 'Fast Food', 'Example Tag 3']
  
  return (
    <StyledListItem selected={selected} alignItems="flex-start" focusRipple={false} onClick={onClick}>
      <CardMedia
        component="img"
        sx={{ width: 151, padding: 2 }}
        image={placeholderImageUrl}
        alt={name + ' image'}
      />
      <Box sx={{ display: 'flex', flexGrow: 1, flexDirection: 'column' }}>
        <CardContent sx={{ flex: '1 0 auto', px: 1, py: 3 }}>
          <Typography component="div" sx={{fontWeight: 'normal'}} variant="h4">
            {name}
          </Typography>
          <Box sx={{my: 1}}>
            <Rating rating={rating} numRatings={numRatings} />
          </Box>
          <Box sx={{my: 1}}>
            <Tags tags={tags} />
          </Box>
          <Typography variant="caption" color="text.secondary" component="div">
            {`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec vitae erat vestibulum, convallis risus sit amet, ultrices nisl.`}
            <span style={{fontWeight: 'bold', marginLeft: 5}}>{'–'}{names[rating]}</span>
          </Typography>
        </CardContent>
      </Box>

    </StyledListItem>
  );
};