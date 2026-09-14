const SpkActivityCard = ({ activityCard, showTime = false, Listclass , Salesdashboard, content}: any) => {
  return (
      <li className={Listclass}>
        <div>
            <div className="font-medium text-[13px]">
              {activityCard?.activityUser}
                {showTime && (
                  <span className={Salesdashboard ===true ? "text-xs activity-time":" text-[11px] text-textmuted  dark:text-textmuted/50 float-end"}>
                    {activityCard?.activityTime}
                  </span>
                )}
            </div>
            <span
            className="block text-textmuted dark:text-textmuted/50"
            dangerouslySetInnerHTML={{ __html: activityCard?.activityDesc }}
            />
            {activityCard?.content}
        </div>
      </li>
  );
};

export default SpkActivityCard;
