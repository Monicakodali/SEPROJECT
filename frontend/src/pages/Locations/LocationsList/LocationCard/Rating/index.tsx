import * as React from 'react';
import { Box, Typography } from '@mui/material';
import { styled } from '@mui/material/styles';

const Star = styled('div')(({theme}) => ({
  display: 'inline-block',
  textAlign: 'center',
  lineHeight: '16px',
  fontSize: 14,
  borderRadius: 4,
  color: '#FFF',
  height: 16,
  width: 16,
  '&:not(:last-child)': {
    marginRight: 3
  }
}))

type RatingProps = {
  rating: number,
  numRatings?: number
}

const star = '★'
const MAX_RATING = 5
const chars = (new Array(MAX_RATING)).fill(star)

export default function Rating({rating, numRatings}: RatingProps) {
  
  return (
    <Box sx={{display: 'flex'}}>
      <Box sx={{mr: 1.5}}>
        {chars.map((c, i) => <Star key={i} sx={{backgroundColor: (i+1) > rating ? 'lightgray' : 'rgb(255,126,66)'}}>{c}</Star>)}
      </Box>
      <Typography color="text.secondary">
        {numRatings}
      </Typography>
    </Box>
  );
};