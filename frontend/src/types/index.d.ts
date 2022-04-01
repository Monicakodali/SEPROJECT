
type Establishment = {
  id: string,
  name: string
  building: string,
  room: string,
  url: string,
  x: number,
  y: number,
  type: 'restaurant' | 'museum' | 'library' | 'building'
}

type Review = {
  Email: string,
  Name: string,
  Est_id: string,
  Est_name: string,
  Review: string,
  Rating: number,
  revTime: string
}

type User = {
  Username: string,
  Email?: string,
  Name: string,
  Password?: string,
  Verified?: string
}