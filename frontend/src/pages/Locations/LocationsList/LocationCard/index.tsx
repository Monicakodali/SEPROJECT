import * as React from 'react';
import { Box, CardContent, Typography, CardMedia, ListItemButton } from '@mui/material';
import { styled } from '@mui/material/styles';
import { Link, Rating, Tags } from '../../../../components'

type LocationsCardProps = {
  data: Diner,
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
      //boxShadow: theme.shadows[3],
      //backgroundColor: 'initial' //modify hover bg color?
    }
  }
})

const StyledLink = styled(Link)(({ theme }) => {
  return {
    fontWeight: 'normal',
    color: 'inherit',
    textDecoration: 'none',
    '&:hover, &:focus': {
      textDecoration: 'underline'
    }
  }
})

const names = ['', 'John', 'Mary', 'Sue', 'Bob', 'Mike']


type InteractiveElementType = {
  name: Diner['Name'],
  id: number | string
  tags: string[],
  rating: number,
  numRatings: number
}

function InteractiveElement({name, id, tags, rating, numRatings}: InteractiveElementType) {
  return (<Box sx={{position: 'absolute', left: 191, top: 24, backgroundColor: 'transparent'}}>
      <StyledLink to={`/est/${id}`} variant="h4">
        {name}
      </StyledLink>
      <Box sx={{my: 1}}>
        <Rating rating={rating} numRatings={numRatings} />
      </Box>
      <Box sx={{my: 1}}>
        <Tags variant="chips" tags={tags} />
      </Box></Box>)
}




export default function LocationsCard({data, selected, onClick}: LocationsCardProps) {

  const { Name: name, Est_Id: id } = data

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
    <Box sx={{position: 'relative'}}>
    <StyledListItem selected={selected} alignItems="flex-start" focusRipple={false} onClick={onClick}>
      <CardMedia
        component="img"
        sx={{ width: 151, padding: 2 }}
        image={placeholderImageUrl}
        alt={name + ' image'}
      />
      <Box sx={{ display: 'flex', flexGrow: 1, flexDirection: 'column' }}>
        <CardContent sx={{ flex: '1 0 auto', px: 1, py: 3 }}>
          <Box sx={{minHeight: 100}}/> 
          <Typography variant="caption" color="text.secondary" component="div">
            {`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec vitae erat vestibulum, convallis risus sit amet, ultrices nisl. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec vitae erat vestibulum, convallis risus sit amet, ultrices nisl. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec vitae erat vestibulum, convallis risus sit amet, ultrices nisl.`}
            <span style={{fontWeight: 'bold', marginLeft: 5}}>{'–'}{names[rating]}</span>
          </Typography>
        </CardContent>
      </Box>
    </StyledListItem>
    <InteractiveElement id={id} name={name} tags={tags} numRatings={numRatings} rating={rating} />
    </Box>
  );
};