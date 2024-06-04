DROP TABLE IF EXISTS merchants;

DO $$
BEGIN
    -- Check and drop the earthdistance extension if it exists
    IF EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'earthdistance') THEN
        DROP EXTENSION earthdistance;
    END IF;
    
    -- Check and drop the cube extension if it exists
    IF EXISTS (SELECT 1 FROM pg_extension WHERE extname = 'cube') THEN
        DROP EXTENSION cube;
    END IF;
    
END $$;