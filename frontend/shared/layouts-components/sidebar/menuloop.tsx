
import { forwardRef, Fragment } from "react";
import { connect } from "react-redux";
import Link from "next/link";
import { ThemeChanger } from "@/shared/redux/action";
import SpkButton from "@/shared/@spk-reusable-components/uielements/spk-button";
import { useLanguage } from "@/shared/i18n/LanguageContext";

const Menuloop = forwardRef(({ MenuItems, toggleSidemenu, local_varaiable, level, HoverToggleInnerMenuFn }: any, ref: any) => {

  const { t } = useLanguage();

  const handleClick = (event: any) => {
    event.preventDefault();
  };

  return (
    <Fragment>
      <Link href="#!" scroll={false} className={`side-menu__item ${MenuItems?.selected ? "active" : ""}`} onMouseOver={(event) => HoverToggleInnerMenuFn(event, MenuItems)}
        onClick={(event) => { event.preventDefault(); toggleSidemenu(event, MenuItems, undefined, true); }} >
        <span className={`${local_varaiable?.dataVerticalStyle == 'doublemenu' ? '' : 'd-none'}`}>
        <span className={`hs-tooltip inline-block [--placement:right] leading-none ${local_varaiable?.dataVerticalStyle == 'doublemenu' ? '' : 'hidden'}`}>
            <SpkButton buttontype="button" customClass="hs-tooltip-toggle  inline-flex justify-center items-center">
                  {MenuItems.icon}
                <span className="hs-tooltip-content hs-tooltip-shown:opacity-100 hs-tooltip-shown:visible opacity-0 transition-opacity inline-block absolute invisible z-10 !py-2 !px-3 !rounded-md bg-black text-xs font-medium text-white shadow-sm dark:bg-black" role="tooltip">
                    {t(MenuItems.title)}
              </span>
            </SpkButton>
          </span>
        </span>
        {local_varaiable?.dataVerticalStyle != "doublemenu" ? MenuItems.icon : ""}

        <span className={`${level == 1 ? "side-menu__label" : ""}`}> {t(MenuItems.title)} {MenuItems.badgetxt ? (<span className={MenuItems.class}> {MenuItems.badgetxt} </span>
        ) : (
          ""
        )}
        </span>
        <i className="ri-arrow-down-s-line side-menu__angle"></i>
      </Link>
      <ul ref={ref} className={`slide-menu child${level}  ${MenuItems.active ? 'double-menu-active' : ''} ${MenuItems?.dirchange === true ? "force-left" : ""} `} style={MenuItems.active ? { display: "block" } : { display: "none" }}>
        {level <= 1 ? <li className="slide side-menu__label1">
          <Link href="#!" scroll={false}>{t(MenuItems.title)}</Link>
        </li> : ""}
        {MenuItems.children.map((firstlevel: any, index: any) =>
          <li className={`${firstlevel.menutitle ? 'slide__category' : ''} ${firstlevel?.type == 'empty' ? 'slide' : ''} ${firstlevel?.type == 'link' ? 'slide' : ''} ${firstlevel?.type == 'sub' ? 'slide has-sub' : ''} ${firstlevel?.active ? 'open' : ''} ${firstlevel?.selected ? 'active' : ''}`} key={index}>
            {firstlevel.type === "link" ?
              <Link href={firstlevel.path} className={`side-menu__item ${firstlevel.selected ? 'active' : ''}`}>{firstlevel.icon}
                <span className=""> {t(firstlevel.title)} {firstlevel.badgetxt ? (<span className={firstlevel.class}> {firstlevel.badgetxt}</span>
                ) : (
                  ""
                )}
                </span>
              </Link>
              : ""}
            {firstlevel.type === "empty" ?
              <Link href="#!" className='side-menu__item' onClick={handleClick}> {firstlevel.icon}<span className=""> {t(firstlevel.title)} {firstlevel.badgetxt ? (<span className={firstlevel.class}> {firstlevel.badgetxt} </span>
              ) : (
                ""
              )}
              </span>
              </Link>
              : ""}
            {firstlevel.type === "sub" ?
              <Menuloop MenuItems={firstlevel} toggleSidemenu={toggleSidemenu} HoverToggleInnerMenuFn={HoverToggleInnerMenuFn} level={level + 1} />
              : ''}

          </li>
        )}

      </ul>
    </Fragment>
  );
})

const mapStateToProps = (state: any) => ({
  local_varaiable: state
});

export default connect(mapStateToProps, { ThemeChanger })(Menuloop);
