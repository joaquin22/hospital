import SpkButton from '@/shared/@spk-reusable-components/uielements/spk-button'
import Link from 'next/link'
import React, { Fragment } from 'react'

const Pageheader = (props: any) => {
  return (
    <Fragment>

{/* <!-- Page Header --> */}
        <div className="flex items-center justify-between page-header-breadcrumb flex-wrap gap-2">
            <div>
                <nav>
                    <ol className="breadcrumb mb-1 flex-wrap">
                        <li className="breadcrumb-item"><Link scroll={false} href="#!">{props.title}</Link></li>
                          {props.subtitle && (
                            <li className="breadcrumb-item"><Link scroll={false} href="#!">{props.subtitle}</Link></li>
                          )}
                      <li className="breadcrumb-item active" aria-current="page">{props.currentpage}</li>
                    </ol>
                </nav>
                <h1 className="page-title font-medium text-lg mb-0">{props.activepage}</h1>
            </div>
            <div className="">
                <SpkButton customClass="ti-btn bg-white dark:bg-bodybg border border-defaultborder !me-2 dark:border-defaultborder/10 btn-wave !my-0">
                    <i className="ri-filter-3-line align-middle me-1 leading-none me-2"></i> Filter
                </SpkButton>
                <SpkButton variant="primary" customClass="ti-btn !border-0 btn-wave !m-0">
                    <i className="ri-share-forward-line me-1"></i> Share
                </SpkButton>
            </div>
        </div>
 {/* <!-- Page Header Close --> */}

    </Fragment>
  )
}

export default Pageheader