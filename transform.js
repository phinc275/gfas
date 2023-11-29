function transform(original) {
    // Log the transform calls
    // log.info(
    //     "Transforming event {Type} from {Stream}",
    //     original.EventType, original.Stream
    // );

    // Ignore some events
    if (original.EventType !== 'USER_ACHIEVEMENT_COMPLETED') {
        return undefined;
    }

    const prefix = 'USER_ACHIEVEMENTS-'
    if (!original.Stream.startsWith(prefix)) {
        return undefined;
    }

    let userId = original.Stream.substring(prefix.length);

    // Create a new event version
    const newEvent = {
        // Copy original data
        ...original.Data,
        user_id: userId
    };

    // Return the new proposed event with modified stream and type
    return {
        Stream: `achievements-${userId}`,
        EventType: original.EventType,
        Data: newEvent,
        Meta: original.Meta
    }
}