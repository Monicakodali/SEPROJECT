
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

type Diner = {
  est_id: string,
  Type: string
  Name: string,
  Building: string,
  Room: string,
  x: number,
  y: number
}

type Review = {
  Review_Id: number,
  Review_user: string,
  Review_est: int,
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