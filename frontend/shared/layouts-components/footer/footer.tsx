
import Link from 'next/link'
import React, { Fragment } from 'react'

const Footer = () => {
  return (
    <Fragment>
      <footer className="footer mt-auto py-4 bg-white dark:bg-bodybg  border-defaultborder dark:border-defaultborder/10 text-center">
        <div className="container">
          <span className="text-textmuted dark:text-textmuted/50"> Copyright © <span className='text-textmuted dark:text-textmuted/50' id="year"> 2025 </span>
            <Link href="#!" scroll={false} className="text-dark font-medium">Xintra</Link>. Designed with <span className="bi bi-heart-fill text-danger"></span> by <Link href="https://spruko.com/" scroll={false} target='_blank'>
              <span className="font-medium text-primary">Spruko</span>
            </Link> All rights reserved </span>
        </div>
      </footer>
    </Fragment>
  )
}

export default Footer