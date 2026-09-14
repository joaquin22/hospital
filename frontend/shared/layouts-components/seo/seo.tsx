
import React from 'react';
import favicon from "../../../public/favicon.ico";
import Head from 'next/head';

const Seo = ({ title }: any) => {
  let i = `Xintra - ${title}`;

  return (
    <>
    <Head>
        <title>{i}</title>
        <link href={favicon.src} rel="icon"></link>
        <meta name="description" content="Xintra - Nextjs Admin &amp; Dashboard Template" />
        <meta name="author" content="Spruko Technologies Private Limited" />
        <meta name="keywords" content="nextjs template typescript, next js, admin panel template, admin, next js typescript, admin template, admin dashboard template, react typescript, dashboard, react next js, nextjs admin template, nextjs bootstrap, admin panel, bootstrap next js" />
      </Head>
    </>
  )
}

export default Seo
 