import type { FC } from 'react'
import { auth } from '../../firebase'

export const UserInfo: FC = () => {
  return (
    <div>
      <img src={auth.currentUser?.photoURL || ''} alt="" />
      <p>{auth.currentUser?.displayName}</p>
    </div>
  )
}
