
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
  FirstName: string,
  LastName: string,
  Address?: string
  Password?: string,
  Verified?: 1 | 0
}