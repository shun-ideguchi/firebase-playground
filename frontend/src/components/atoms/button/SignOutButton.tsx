import type { FC } from 'react'
import { auth } from '../../../firebase'

export const SignOutButton: FC = () => {
  return (
    <button onClick={() => auth.signOut()}>
      <p>サインアウト</p>
    </button>
  )
}
