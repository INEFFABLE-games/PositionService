
CREATE FUNCTION notify_trigger_open_close() RETURNS trigger as $trigger$
DECLARE
rec RECORD;
BEGIN

CASE TG_OP
        WHEN 'INSERT' THEN
            rec := NEW;
            perform pg_notify('db_notifications_open', row_to_json(NEW)::text);
        WHEN 'UPDATE' THEN
            rec := NEW;
            perform pg_notify('db_notifications_close', row_to_json(NEW)::text);
ELSE
            RAISE EXCEPTION 'Unknown TG_OP: "%". Should not occur!', TG_OP;
END CASE;
return rec;
end;


$trigger$ LANGUAGE plpgsql